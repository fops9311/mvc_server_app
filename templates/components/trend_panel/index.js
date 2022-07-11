var labels = ["Eve", "Cain", "Seth", "Enos", "Noam", "Abel", "Awan", "Enoch", "Cain"]

var parents = ["", "Eve", "Eve", "Seth", "Seth", "Eve", "Eve", "Awan", "Awan"]

var data = [{

      type: "treemap",

      branchvalues: "total",

      labels: labels,

      parents: parents,

      domain: {x: [-1, 1]},

      values: [65, 14, 12, 10, 2, 6, 6, 2, 4],

      textinfo: "label+value+percent parent+percent entry",

      outsidetextfont: {"size": 40, "color": "#377eb8"},

      marker: {"line": {"width": 1}},

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
