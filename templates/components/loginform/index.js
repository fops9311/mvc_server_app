document.addEventListener('submit', (e) => {
    // Store reference to form to make later code easier to read
    const form = e.target;
    if (e.target.id == "loginform"){
      const data = new URLSearchParams();
      for (const pair of new FormData(form)) {
        data.append(pair[0], pair[1]);
      }
      var login = document.getElementById("user_id").value;
      var password = document.getElementById("password").value;
      // Post data using the Fetch API
      fetch(form.action, {
        method: form.method,
        body: data,
      }).then(function(response) {
        return response.text().then(function(text) {
            console.log("try login as "+login+" pass: "+password)
            if (text==="Ok"){
                localStorage.setItem("login",login)
                localStorage.setItem("password",password)
                alert("Success ".concat(login," ",password))
                window.location = "/{{ .version}}/pages"
            }
            else{
              alert("Fail")
            }
        });
    });
      // Prevent the default form submit
      e.preventDefault();
    }
});