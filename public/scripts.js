document.addEventListener("DOMContentLoaded", function () {
  document.body.addEventListener("htmx:beforeSwap", function (event) {
    if (event.detail.xhr.status === 422) {
      event.detail.shouldSwap = true;
      event.detail.isError = false;
    }

    // if (event.detail.xhr.status >= 300 && event.detail.xhr.status < 400) {
    //   // Get the 'Location' header from the response
    //   const location = event.detail.xhr.getResponseHeader("Location");
    //   if (location) {
    //     // Perform client-side redirect
    //     window.location.href = location;
    //   }
    // }
  });
});
