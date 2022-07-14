
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

//Plotly.newPlot('myDiv', data, layout, config)

var myPlot = document.getElementById('myDiv')

myPlot.on('plotly_click', function(data){
  localStorage.setItem('current_obj', data.points[0].label);
  console.log(localStorage.getItem('current_obj') );
});