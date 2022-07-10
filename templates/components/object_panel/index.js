function refresh_obj_list(){
    object_list = document.getElementById("object_list");
    fetch("/v1/users/fops9311@yandex.ru/objects", {
        method: "GET",
    }).then(function(response) {
        return response.text().then(function(text) {
            object_list.innerHTML = text;
        });
      });
}