begin;

drop table settings;
drop table page_html;
drop table page;

drop function cms_set_updated_at;
drop function cms_page_path_uniqueness;

drop type page_type;

commit;