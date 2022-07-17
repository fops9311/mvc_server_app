asyncCallbacks.push(refresh_trn)
var trn_created = false
function refresh_trn(objects){
    if (!trn_created){
    Plotly.newPlot(
        'myTrend', 
        transformObjectsTraces(objects),
        trendlayout,
        {responsive: true},
    );
    trn_created = false
    }
}
var trendlayout = {
    hovermode: "x unified",
    showlegend: false,
    legend: {
      orientation: "h",
      x: 0,
      xanchor: 'left',
      yanchor: 'top',
      y: 1},
    title: 'Time Series with Rangeslider',
    xaxis: {
      autorange: true,
      rangeselector: {buttons: [
          {
            count: 1,
            label: '1m',
            step: 'month',
            stepmode: 'backward'
          },
          {
            count: 6,
            label: '6m',
            step: 'month',
            stepmode: 'backward'
          },
          {step: 'all'}
        ]},
      rangeslider: {range: ['2022-07-06', '2022-07-20']},
      type: 'date'
    },
    yaxis: {
      autorange: true,
      fixedrange: false,
      type: 'linear'
    }
  };
  
console.log("testting tranformSamplesTrace")
console.log(tranformSamplesTrace(
    [
    {
        "Timestamp": "2022-07-10T09:12:58.09556692+03:00",
        "Value": 0.04
    },
    {
        "Timestamp": "2022-07-10T09:13:05.909189508+03:00",
        "Value": 0.07
    }
    ],
    "test",
    ))

function tranformSamplesTrace(samples,id){
    var trace = {
        type:"scatter",
        mode:"lines+markers",
        name:transformIds1(id),
        x:[],
        y:[],
        line: {
            //color: '#17BECF',
            shape: 'linear'
        }
    }
    samples.forEach(s=>{
        trace.x.push(s.Timestamp.replace("T"," "))
        trace.y.push(s.Value)
    })
    
    return trace
}

console.log("testting transformObjectsTraces")
console.log(transformObjectsTraces({
	"dsd/d/d": {
		"Id": "dsd/d/d",
		"Samples": [
			{
				"Timestamp": "2022-07-09T20:45:57.840348668+03:00",
				"Value": 42.14
			}
		],
		"LastSample": {
			"Timestamp": "0001-01-01T00:00:00Z",
			"Value": 0
		}
	},
	"fops9311/factory1": {
		"Id": "fops9311/factory1",
		"Samples": [
			{
				"Timestamp": "2022-07-09T20:48:40.759329754+03:00",
				"Value": 0.02
			},
			{
				"Timestamp": "2022-07-09T20:48:54.109789467+03:00",
				"Value": 0.04
			}
		],
		"LastSample": {
			"Timestamp": "0001-01-01T00:00:00Z",
			"Value": 0
		}
	}
}))
function transformObjectsTraces(objects) {
    var traces = []
    for (const [key, value] of Object.entries(objects).sort((a, b) => a[1].Id > b[1].Id)) {
        console.log(`${key}: ${value}`);
        value.Samples.push(value.LastSample)
        traces.push(tranformSamplesTrace(value.Samples,value.Id))
      }
    return traces
}
console.log("testing transformIds1...")
console.log(transformIds1("test/t/test/23"))
function transformIds1(id){
    
    var parts = id.split("/")
    parts.shift()
    var result = []
    parts.forEach((part,i)=>{
        if (i<(parts.length-1)){
            result.push(part.concat("/"))
        }else{
            result.push("<b>".concat(part.concat("</b>")))
        }
    })
    return result.join("")
}