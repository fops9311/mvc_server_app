document.addEventListener('submit', (e) => {
    // Store reference to form to make later code easier to read
    const form = e.target;
    if (e.target.id == "new_sample"){
      var object_id =""
      var sample_value =""
      for (const pair of new FormData(form)) {
        switch (pair[0]){
            case "object_id": object_id = pair[1];break;
            case "sample_value": sample_value = pair[1]; break;
        }
      }
      // Post data using the Fetch API
      fetch("/{{ .version}}/users/"+localStorage.getItem("login")+"/objects/"+encodeURIComponent(object_id)+"/now?"+ new URLSearchParams({
        password: localStorage.getItem("password"),
        sample_value: sample_value,
        }), {
        method: form.method,
      });
      // Prevent the default form submit
      e.preventDefault();
    }
});