var userObjects
var asyncCallbacks = []

var appUser = localStorage.getItem("login")

var pageDisplayCallbacks = []
var REFRESH_RATE = 5000
var app = {
    activePage: 0,
    auth: false,

}
const loginformPage         = 0
const registerPage          = 1
const object_panelPage      = 2
const trend_panelPage       = 3
const controls_panelPage    = 4

function isActivePage(pageId){
    return app.activePage == pageId 
}
function displayAsPage(elem_id,page_id){
    f = function(){
        if (isActivePage(page_id)){
            document.getElementById(elem_id).style.display=""
            //console.log(page_id.concat(" display=true"))
        }else{
            document.getElementById(elem_id).style.display="none"
            //console.log(page_id.concat(" display=false"))
        }
    }
    pageDisplayCallbacks.push(f)

    f()
}
function displayPageById(page_id){
    app.activePage = page_id
    pageDisplayCallbacks.forEach(f=>{f()})
}

function ifPageDisplayed(ifTrue,ifFalse,page_id){
    var t = function(){
        if (isActivePage(page_id)){
            ifTrue()
        }
    }
    var f = function(){
        if (!isActivePage(page_id)){
            ifFalse()
        }
    }
    pageDisplayCallbacks.push(t)
    pageDisplayCallbacks.push(f)
    if (isActivePage(page_id)){
        t()
    }else{
        f()
    }
}
    

setInterval(refresh_userObjects,REFRESH_RATE)
function refresh_userObjects() {
    fetch("/v1/users/"+localStorage.getItem("login")+"/objects?"+ new URLSearchParams({
        password: localStorage.getItem("password"),
        subdir: localStorage.getItem("subdir"),
    }), {
        method: "GET",
    }).then(function(response) {
        return response.text().then(function(text) {
            userObjects = JSON.parse(text)
            console.log(`[refresh_userObjects][network][result] userObjects.length:${userObjects.length};`);
            asyncCallbacks.forEach(
                c=>{
                    if (c!=null){
                        c(userObjects)
                    }
                })
        });
    });
}

var refresh_object_list_callbacks = []

function refresh_object_list(){
    object_list = document.getElementById("object_list");

    fetch("/v1/users/"+localStorage.getItem("login")+"/objects?"+ new URLSearchParams({
        password: localStorage.getItem("password"),
        subdir: "",
    }), {
        method: "GET",
    }).then(function(response) {
        return response.text().then(function(text) {
            userObjects = JSON.parse(text)
            console.log(`[refresh_object_list][network][result] total_objects.length:${userObjects.length};`);
            refresh_object_list_callbacks.forEach(
                c=>{
                    if (c!=null){
                        c(userObjects)
                    }
                })
        });
    });

}