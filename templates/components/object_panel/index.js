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
var labels = ["fops9311@yandex.ru", "fops9311@yandex.ru/v1", "fops9311@yandex.ru/v2", "fops9311@yandex.ru/v3",  "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v5", "fops9311@yandex.ru/v3/v4/1","fops9311@yandex.ru/v3/v4/2", "fops9311@yandex.ru/v3/v4/311"]

var parents = ["", "fops9311@yandex.ru", "fops9311@yandex.ru", "fops9311@yandex.ru", "fops9311@yandex.ru/v3", "fops9311@yandex.ru/v3", "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v4", "fops9311@yandex.ru/v3/v4"]

var data = [{

      type: "treemap",

      branchvalues: "remainder",

      count: 'branches+leaves',

      maxdepth: 2,

      labels: labels,

      parents: parents,

      domain: {x: [0, 1]},

      textinfo: "label",

      labeltemplate:"'%{customdata[0]}'",

      texttemplate: "Name: %{customdata[0]}<br>Last value: %{customdata[1]}", 

      textposition: 'middle center',

      textfont: {size:16},

      outsidetextfont: {"size": 40, "color": "#377eb8"},

      marker: {"line": {"width": 1}},

      customdata: [
        {0:'root',1:4},
        {0:'v1',1:42},
        {0:'v2',1:41},
        {0:'v3',1:44},
        {0:'v4',1:43},
        {0:'v5',1:24},
        {0:'1',1:14},
        {0:'2',1:422},
        {0:'3',1:490}
      ],


      //hovertemplate: '%{customdata[0]} last value %{customdata[1]}',

      pathbar: {"visible": true}

    }];

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


var config = {responsive: true}

Plotly.newPlot('myDiv', data, layout, config)
