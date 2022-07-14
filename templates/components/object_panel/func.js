var labels1 = ["fops9311@yandex.ru", "fops9311@yandex.ru/v1", "fops9311@yandex.ru/v2", "fops9311@yandex.ru/v3",  "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v5", "fops9311@yandex.ru/v3/v4/1","fops9311@yandex.ru/v3/v4/2", "fops9311@yandex.ru/v3/v4/311"]

var parents1 = ["", "fops9311@yandex.ru", "fops9311@yandex.ru", "fops9311@yandex.ru", "fops9311@yandex.ru/v3", "fops9311@yandex.ru/v3", "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v4"]
var appUser = "fops9311@yandex.ru"
var layout = {
  margin: {

    l: 5,

    r: 5,

    b: 5,

    t: 50,

    pad: 4

  },

  annotations: [{

    showarrow: false,

    text: "branchvalues: <b>total</b>",

    x: 0.5,

    xanchor: "center",

    y: 1.05,

    yanchor: "bottom"

    }]}
    
function treedata(l,p,customdata){
  l.push(appUser)
  p.push("")
  //customdata=customdata.sort((a, b) => a.length - b.length)
  customdata.push({0:"root",1:0})
  return [{

    type: "treemap",

    branchvalues: "remainder",

    count: 'branches+leaves',

    maxdepth: 2,

    labels: l,

    parents: p,

    domain: {x: [0, 1]},

    textinfo: "label",

    labeltemplate:"'%{customdata[0]}'",

    texttemplate: "Name: %{customdata[0]}<br>Last value: %{customdata[1]}", 

    textposition: 'middle center',

    textfont: {size:16},

    outsidetextfont: {"size": 40, "color": "#377eb8"},

    marker: {"line": {"width": 1}},

    customdata: customdata,


    //hovertemplate: '%{customdata[0]} last value %{customdata[1]}',

    pathbar: {"visible": true}

  }];
}

function refresh_obj_list(){
    object_list = document.getElementById("object_list");
    fetch("/v1/users/fops9311@yandex.ru/objects", {
        method: "GET",
    }).then(function(response) {
        return response.text().then(function(text) {
          Objects = JSON.parse(text)
          console.log(Objects)
          lp = transformIds(getObjectsIds(Objects))
          tdata = treedata(lp.labels,lp.parents,getObjectsCustomdata(Objects))
          Plotly.newPlot('myDiv', tdata, layout, {responsive: true});
          console.log(tdata);
        });
      });
}



console.log("test transformObjects")
res = transformIds(["j1/3","2/j1","j5","j6/2","j1","j1/3","2","j1","j1","j1"])
//console.log(res.labels)
//console.log(res.parents)

function getObjectsCustomdata(objects){
  var result = []
  Object.entries(objects).sort((a, b) => a[1].Id.length - b[1].Id.length).map(o => {
    //console.log(o[1])
    //return
    var parts = o[1].Id.split("/")
    parts.shift()
    //console.log(parts)
    result.push({
      0:parts.join("/"),
      1:o[1].Samples.pop().Value,
    }) 
  })
  return result
}

function getObjectsIds(objects){
  var result = []
  objects.forEach((o)=>{
    result.push(o.Id)
  })
  return result
}

function transformIds(Id){
  var result = {labels:[],parents:[]}
  result.labels=Id.sort((a, b) => a.length - b.length)
  result.labels.every((v)=>{
    parentCandidates(v).every((candidate)=>{
      if (result.labels.includes(candidate) ){
          result.parents.push(candidate)
        return false
      }
      if (candidate===""){
          result.parents.push(appUser)
        return false
      }
      return true
    })
    return true
    })
    console.log(result.labels)
    console.log(result.parents)
    return result
}
function parentCandidates(label){
  dirs = label.split("/");
  if (dirs.length<2){
    return [""]
  }
  result = [];
  for (var i = 0; i < (dirs.length-1); i++) {
    result[i]= dirs.slice(0,(dirs.length-1)-i).join("/")
  }
  result.push("")
  return result
}