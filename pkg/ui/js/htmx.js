import 'htmx.org';
import _hyperscript from 'hyperscript.org';

_hyperscript.browserInit();

window.htmx = require('htmx.org');

document.addEventListener('DOMContentLoaded', function() {
    document.body.addEventListener('htmx:beforeSwap', function (evt) {
        if(evt.detail.xhr.status === 404){
            alert("Error: Not Found (404)");
        } else if (evt.detail.xhr.status === 422) {
            evt.detail.shouldSwap = true;
            evt.detail.isError = false;
        }
    });
});