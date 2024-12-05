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
    meta_robots                 varchar(50),
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
    tags                        text[]                      not null default '{}',
    categories                  text[]                      not null default '{}',
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
    meta_robots                 varchar(50),
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
    -- Common search vector components
    NEW.search_vector :=
        setweight(to_tsvector('english', NEW.title), 'A') ||
        setweight(to_tsvector('english', array_to_string(NEW.tags, ' ')), 'B') ||
        setweight(to_tsvector('english', NEW.meta_description), 'C');

    -- Common full text components
    NEW.full_text := NEW.title || '. ' || array_to_string(NEW.tags, ' ') || '. ' || NEW.meta_description;

    if TG_TABLE_NAME = 'page_html' then
        -- Process HTML content for page_html table
        NEW.full_text := NEW.full_text || '. ' ||
            coalesce(regexp_replace(NEW.html, '<[^>]*>|[\r\n]+|\s{2,}', '', 'g'), '');

        NEW.search_vector := NEW.search_vector ||
            setweight(
                to_tsvector(
                    'english',
                    coalesce(regexp_replace(NEW.html, '<[^>]*>|[\r\n]+|\s{2,}', '', 'g'), '')
                ),
                'D'
            );
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

commit;