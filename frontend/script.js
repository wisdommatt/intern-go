const labels = document.querySelectorAll('.form-control label')

labels.forEach(label => {
    label.innerHTML = label.innerText
    .split('')
    .map((letter, idx) => `<span style="transition-delay:${idx *50}ms"> ${letter}</span>`)
    .join('')
});

var GRAPHQL_URL = "https://api.interngo.works/graphql/query";
if (window.location.host.includes("localhost")) {
    GRAPHQL_URL = "http://localhost:8080/graphql/query";
}

function makeGraphqlRequest(query, callBacks) {
    var url = GRAPHQL_URL;
    $.ajax({
      url: url,
      method: "POST",
      data: JSON.stringify(query),
      contentType: "application/json",
      success: function (response) {
        if (response.errors) {
          if (callBacks.error) {
            callBacks.error(response.errors[0].message);
            return;
          }
          defaultGraphqlErrorHandler(response.errors[0].message);
        } else {
          callBacks.success(response.data);
        }
      },
      error: function (req, status, err) {
        console.log("request error: ", req, status, err, JSON.stringify(query), url);
        Swal.fire("Error", "network connection error", "error");
      }
    });
  }

  function defaultGraphqlErrorHandler(errMsg) {
    Swal.fire("Error", errMsg, "error");
  }

  function getFormValues(form) {
    var formData = new FormData(form);
    var formBody = {};
    for (var data of formData.entries()) {
      formBody[data[0]] = data[1];
    }
    return formBody;
  }