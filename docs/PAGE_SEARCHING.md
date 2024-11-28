# Page searching

When creating new page types, if you want them searchable you will need to update the search function that
populates the fields `search_vector` and `full_text`.

See the function in the db `cms_set_page_search_vector`.

you will need to add another trigger to your new page table to invoke the procedure ie:

```sql
create trigger set_search_index
    before insert or update
    on <my_inherited_page_table>
    for each row
    execute procedure cms_set_page_search_vector();
```