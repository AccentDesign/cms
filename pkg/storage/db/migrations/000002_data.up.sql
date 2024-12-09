begin;

insert into settings (meta_og_site_name, meta_robots)
values
    (
     'cms',
     'index, follow'
    );

insert into page_html (path, title, featured_image, meta_description, html)
values
    (
     '',
     'home',
     'https://placehold.co/600/6b7280/f9fafb?text=home&font=open-sans',
     'description for the home page',
     '<div class="space-y-6"><div class="h-48 bg-gray-100 rounded-lg"></div><div class="space-y-3"><p>Some crappy text on the home page.</p><div class="h-4 bg-gray-100 rounded w-2/3"></div><div class="h-4 bg-gray-100 rounded w-1/2"></div><div class="h-4 bg-gray-100 rounded w-full"></div></div></div>'
    );

insert into page (path, title, no_cache, meta_description, meta_robots, page_type, is_in_sitemap, is_searchable)
values
    (
     'about',
     'about',
     false,
     'description for the about page',
     null,
     'listing',
     true,
     true
    ),
    (
    'search',
    'search',
    true,
    'description for the search page',
    'noindex, follow',
    'search',
    false,
    false
    );

insert into page_html (path, title, featured_image, tags, categories, meta_description, html)
values
    (
     'about.dave',
     'dave',
     'https://placehold.co/600/6b7280/f9fafb?text=dave&font=open-sans',
     '{"Owner"}',
     '{"Team"}',
     'description for the dave page',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><img _="init set me.src to window.pageData.featured_image" alt="Featured Image" class="h-32 w-32 bg-gray-100 rounded-full"><div class="space-y-2"><div class="h-4 bg-gray-100 rounded w-1/2"></div><div class="h-4 bg-gray-100 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im dave a strapping six footer from the rough end of the trench.</p><div class="h-4 bg-gray-100 rounded w-full"></div><div class="h-4 bg-gray-100 rounded w-5/6"></div><div class="h-4 bg-gray-100 rounded w-4/6"></div></div></div>'
    ),
    (
     'about.karen',
     'karen',
     'https://placehold.co/600/6b7280/f9fafb?text=karen&font=open-sans',
     '{"Manager"}',
     '{"Team"}',
     'description for the karen page',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><img _="init set me.src to window.pageData.featured_image" alt="Featured Image" class="h-32 w-32 bg-gray-100 rounded-full"><div class="space-y-2"><div class="h-4 bg-gray-100 rounded w-1/2"></div><div class="h-4 bg-gray-100 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im karen, daves better half.</p><div class="h-4 bg-gray-100 rounded w-full"></div><div class="h-4 bg-gray-100 rounded w-5/6"></div><div class="h-4 bg-gray-100 rounded w-4/6"></div></div></div>'
    ),
    (
     'about.geoff',
     'geoff',
     'https://placehold.co/600/6b7280/f9fafb?text=geoff&font=open-sans',
     '{"Owner"}',
     '{"Team"}',
     'description for the geoff page',
     '<div class="space-y-6"><div class="flex items-center space-x-4"><img _="init set me.src to window.pageData.featured_image" alt="Featured Image" class="h-32 w-32 bg-gray-100 rounded-full"><div class="space-y-2"><div class="h-4 bg-gray-100 rounded w-1/2"></div><div class="h-4 bg-gray-100 rounded w-1/3"></div></div></div><div class="space-y-3"><p>Hello im geoff and I love a good factory.</p><div class="h-4 bg-gray-100 rounded w-full"></div><div class="h-4 bg-gray-100 rounded w-5/6"></div><div class="h-4 bg-gray-100 rounded w-4/6"></div></div></div>'
    );

commit;