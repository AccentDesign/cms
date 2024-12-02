begin;

-- extension

create extension if not exists ltree;

-- types

create type page_type as enum (
    'general',
    'listing',
    'search'
);

-- functions

create or replace function cms_set_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = clock_timestamp();
return NEW;
end;
$$ language plpgsql;

-- settings

create table settings
(
    id                          serial                      not null primary key check (id = 1),
    meta_description            varchar(320),
    meta_og_site_name           varchar(320),
    meta_og_title               varchar(320),
    meta_og_description         varchar(320),
    meta_og_url                 varchar(320),
    meta_og_type                varchar(320),
    meta_og_image               varchar(320),
    meta_og_image_secure_url    varchar(320),
    meta_og_image_width         varchar(320),
    meta_og_image_height        varchar(320),
    meta_article_publisher      varchar(320),
    meta_article_section        varchar(320),
    meta_article_tag            varchar(320),
    meta_twitter_card           varchar(320),
    meta_twitter_image          varchar(320),
    meta_twitter_site           varchar(320),
    created_at                  timestamp                   not null default clock_timestamp(),
    updated_at                  timestamp                   not null default clock_timestamp()
);

create trigger set_updated_at
    before update
    on settings
    for each row
    execute procedure cms_set_updated_at();

-- page

create table page
(
    id                          serial                      not null primary key,
    path                        ltree                       not null,
    level                       int                         generated always as (nlevel(path)) stored,
    url                         text                        generated always as ('/' || replace(path::text, '.', '/')) stored,
    page_type                   page_type                   not null default 'general',
    title                       varchar(160)                not null,
    is_in_sitemap               boolean                     not null default true,
    is_searchable               boolean                     not null default true,
    search_vector               tsvector                    not null default to_tsvector(''),
    full_text                   text                        not null default '',
    no_cache                    boolean                     not null default false,
    created_at                  timestamp                   not null default clock_timestamp(),
    updated_at                  timestamp                   not null default clock_timestamp(),
    meta_description            varchar(320),
    meta_og_site_name           varchar(320),
    meta_og_title               varchar(320),
    meta_og_description         varchar(320),
    meta_og_url                 varchar(320),
    meta_og_type                varchar(320),
    meta_og_image               varchar(320),
    meta_og_image_secure_url    varchar(320),
    meta_og_image_width         varchar(320),
    meta_og_image_height        varchar(320),
    meta_article_publisher      varchar(320),
    meta_article_section        varchar(320),
    meta_article_tag            varchar(320),
    meta_twitter_card           varchar(320),
    meta_twitter_image          varchar(320),
    meta_twitter_site           varchar(320),
    check (
        (page_type != 'general' and tableoid = 'page'::regclass) or
        (page_type = 'general' and tableoid != 'page'::regclass)
    )
);

create index page_path_gist_idx on page using gist (path);
create index page_search_vector_idx ON page USING GIN (search_vector);

create or replace function cms_page_path_uniqueness()
    returns trigger as
$$
begin
    if exists (select 1 from page where path = NEW.path and id <> NEW.id) then
        raise exception 'Path "%" is already in use.', NEW.path;
end if;
return NEW;
end;
$$ language plpgsql;

create or replace function cms_set_page_search_vector()
    returns trigger as
$$
begin
    if TG_TABLE_NAME = 'page' then
        NEW.search_vector = setweight(to_tsvector('english', NEW.title), 'A') || setweight( to_tsvector('english', NEW.meta_description), 'B');
        NEW.full_text = NEW.title || '. ' || NEW.meta_description;

    elseif TG_TABLE_NAME = 'page_html' then
        NEW.search_vector = setweight(to_tsvector('english', NEW.title), 'A') || setweight( to_tsvector('english', NEW.meta_description), 'B')  || setweight( to_tsvector('english', coalesce(regexp_replace(NEW.html, '<[^>]*>|[\r\n]+|\s{2,}', '', 'g'), '')), 'C');
        NEW.full_text = NEW.title || '. ' || NEW.meta_description || '. ' || coalesce(regexp_replace(NEW.html, '<[^>]*>|[\r\n]+|\s{2,}', '', 'g'), '');

end if;
return NEW;
end;
$$ language plpgsql;

create trigger set_updated_at
    before update
    on page
    for each row
    execute procedure cms_set_updated_at();

create trigger set_search_index
    before insert or update
    on page
    for each row
    execute procedure cms_set_page_search_vector();

create trigger check_path_uniqueness
    before insert or update
    on page
    for each row
    execute procedure cms_page_path_uniqueness();

-- page_html

create table page_html
(
    html                        text                        not null
) inherits (page);

create trigger set_updated_at
    before update
    on page_html
    for each row
    execute procedure cms_set_updated_at();

create trigger set_search_index
    before insert or update
    on page_html
    for each row
    execute procedure cms_set_page_search_vector();

create trigger check_path_uniqueness
    before insert or update
    on page_html
    for each row
    execute procedure cms_page_path_uniqueness();

-- example data

insert into settings (meta_og_site_name)
values
    ('localhost');

insert into page_html (path, title, meta_description, meta_og_image, html)
values
    (
     '',
     'home',
     'description for the home page',
     'https://placehold.co/600/e5e7eb/ffffff?text=home&font=open-sans',
     '<div class="space-y-6"><div class="h-48 bg-gray-200 rounded-lg"></div><div class="space-y-3"><p>Some crappy text on the home page.</p><div class="h-4 bg-gray-200 rounded w-2/3"></div><div class="h-4 bg-gray-200 rounded w-1/2"></div><div class="h-4 bg-gray-200 rounded w-full"></div></div></div>'
    );

insert into page (path, title, no_cache, meta_description, meta_og_image, page_type, is_in_sitemap, is_searchable)
values
    (
     'about',
     'about',
     false,
     'description for the about page',
     'https://placehold.co/600/e5e7eb/ffffff?text=about&font=open-sans',
     'listing',
     true,
     true
    ),
    (
    'search',
    'search',
     true,
    'description for the search page',
    'https://placehold.co/600/e5e7eb/ffffff?text=search&font=open-sans',
    'search',
    false,
    false
    );

insert into page_html (path, title, meta_description, meta_og_image, html)
values
    (
     'about.dave',
     'dave',
     'description for the dave page',
     'https://placehold.co/600/e5e7eb/ffffff?text=dave&font=open-sans',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><div class="h-32 w-32 bg-gray-200 rounded-full"></div><div class="space-y-2"><div class="h-4 bg-gray-200 rounded w-1/2"></div><div class="h-4 bg-gray-200 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im dave a strapping six footer from the rough end of the trench.</p><div class="h-4 bg-gray-200 rounded w-full"></div><div class="h-4 bg-gray-200 rounded w-5/6"></div><div class="h-4 bg-gray-200 rounded w-4/6"></div></div></div>'
    ),
    (
     'about.karen',
     'karen',
     'description for the karen page',
     'https://placehold.co/600/e5e7eb/ffffff?text=karen&font=open-sans',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><div class="h-32 w-32 bg-gray-200 rounded-full"></div><div class="space-y-2"><div class="h-4 bg-gray-200 rounded w-1/2"></div><div class="h-4 bg-gray-200 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im karen, daves better half.</p><div class="h-4 bg-gray-200 rounded w-full"></div><div class="h-4 bg-gray-200 rounded w-5/6"></div><div class="h-4 bg-gray-200 rounded w-4/6"></div></div></div>'
    ),
    (
     'about.geoff',
     'geoff',
     'description for the geoff page',
     'https://placehold.co/600/e5e7eb/ffffff?text=geoff&font=open-sans',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><div class="h-32 w-32 bg-gray-200 rounded-full"></div><div class="space-y-2"><div class="h-4 bg-gray-200 rounded w-1/2"></div><div class="h-4 bg-gray-200 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im geoff and I love a good factory.</p><div class="h-4 bg-gray-200 rounded w-full"></div><div class="h-4 bg-gray-200 rounded w-5/6"></div><div class="h-4 bg-gray-200 rounded w-4/6"></div></div></div>'
    );

commit;