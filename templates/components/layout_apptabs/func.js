function drawDirSelect(objects){

    var select = document.getElementById("dirSelect")
    select.innerHTML =""
    select.innerHTML += "<option value=\"\" "+is_selected(()=>{localStorage.getItem("subdir")===""})+">"+appUser+"</option>"

    for (const [key, value] of Object.entries(objects).sort((a, b) => a[1].Id > b[1].Id)) {
      if (is_brunch(value.Id)){
        select.innerHTML += "<option value=\""+cutUser_id(value.Id)+"\" "+is_selected(()=>{return localStorage.getItem("subdir")===cutUser_id(value.Id)})+">"+cutUser_id(value.Id)+"</option>"
      }
      }
    console.log(`[drawDirSelect][result] DONE`);
    function is_selected(cond){
      if (cond()){
        return "selected"
      }
      return ""
    }
    function is_brunch(id){
      for (const [key, value] of Object.entries(objects).sort((a, b) => a[1].Id > b[1].Id)) {
        if (value.Id.includes(id) && value.Id!=id){
          return true
        }
      }
      return false
    }
  }
  function dirSelectChanged(){
    var e = document.getElementById("dirSelect");
    var value = e.options[e.selectedIndex].value;
    //var text = e.options[e.selectedIndex].text;
    localStorage.setItem("subdir",value)
  }

function cutUser_id(id){
  var parts = id.split("/")
  parts.shift()
  var result = []
  parts.forEach((part,i)=>{
      if (i<(parts.length-1)){
          result.push(part.concat("/"))
      }else{
          result.push(part)
      }
  })
  return result.join("")
}