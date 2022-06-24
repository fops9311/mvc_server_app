# mvc_server_app
...
user_post_path  GET     /users/:user_id/posts           HelloWeb.PostController :index
user_post_path  GET     /users/:user_id/posts/:id/edit  HelloWeb.PostController :edit
user_post_path  GET     /users/:user_id/posts/new       HelloWeb.PostController :new
user_post_path  GET     /users/:user_id/posts/:id       HelloWeb.PostController :show
user_post_path  POST    /users/:user_id/posts           HelloWeb.PostController :create
user_post_path  PATCH   /users/:user_id/posts/:id       HelloWeb.PostController :update
                PUT     /users/:user_id/posts/:id       HelloWeb.PostController :update
user_post_path  DELETE  /users/:user_id/posts/:id       HelloWeb.PostController :delete
...