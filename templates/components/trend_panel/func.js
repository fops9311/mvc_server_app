function customDateFormat(d){
return d.toISOString().replace("T"," ").replace("Z"," ")
}
console.log("testing date format_______________________________")
console.log(customDateFormat(new Date(Date.now() - 600000)))
console.log(customDateFormat(new Date(Date.now() + 600000)))
console.log(customDateFormat((function(d){ d.setMinutes(d.getMinutes()+10); return d})(new Date)))
console.log("END_______________________________________________")


asyncCallbacks.push(refresh_trn)
var trn_created = false
function refresh_trn(objects){
    if (!trn_created){
    Plotly.newPlot(
        'myTrend', 
        transformObjectsTraces(objects),
        trendlayout(),
        {responsive: true},
    );
    trn_created = true
    }else{
      var layout = trendlayout()
      var traces = transformObjectsTraces(objects)
      Plotly.react('myTrend',
      {
          'data':traces,
      }
      )
      Plotly.update('myTrend',
      {
          'data':traces,
      },
      {
        'margin.l':layout.margin.l,
        'margin.r':layout.margin.r,
        'margin.b':layout.margin.b,
        'margin.t':layout.margin.t,
        'margin.pad':layout.margin.pad,

        'showlegend':layout.showlegend,
        'legend.orientation':layout.legend.orientation,
        'legend.x':layout.legend.x,
        'legend.y':layout.legend.y,
        'legend.xanchor':layout.legend.xanchor,
        'legend.yanchor':layout.legend.yanchor,
        'xaxis.range':layout.xaxis.range,
        'xaxis.type':layout.xaxis.type,
        'xaxis.autorange':layout.xaxis.autorange,
        'hovermode':layout.hovermode,
      }
      )
     // Plotly.relayout('myTrend',
      //{
        //  'xaxis.range':layout.xaxis.range,
      //})
      console.log(`[trend][react][update] DONE`)
    }
}
var trendlayout = ()=>{return {
  margin: {
      l: 50,
      r: 2,
      b: 50,
      t: 2,
      pad: 15
  },
    hovermode: "x unified",
    showlegend: true,
    legend: {
      orientation: "h",
      x: 0,
      xanchor: 'left',
      yanchor: 'top',
      y: 1},
    xaxis: {
      autorange: false,
      //rangeslider: {range: [customDateFormat(new Date(Date.now() - 600000)), customDateFormat(new Date(Date.now() + 600000))]},
      range: [customDateFormat(new Date(Date.now() - 6000000 + 10800000)), customDateFormat(new Date(Date.now() + 10800000 + 30000))],
      type: 'date'
    },
    yaxis: {
      autorange: true,
      //range:[0,10],
      fixedrange: false,
      type: 'linear'
    }
  }};
  
console.log("testting tranformSamplesTrace_____________________________")
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
console.log("END_______________________________________________")

function tranformSamplesTrace(samples,id){

  const date = new Date();
  const offset = date.getTimezoneOffset();
  //console.log(offset);    // -300 timezone offset in minutes

    var trace = {
        type:"scatter",
        mode:"lines",
        name:transformIds1(id),
        x:[],
        y:[],
        line: {
            //color: '#17BECF',
            shape: 'linear'
        }
    }
    samples.sort((a, b) => a.Timestamp > b.Timestamp).forEach(s=>{
        trace.x.push(customDateFormat(new Date(Date.now() - Date.now() + s.Timestamp/1000000 - offset*1000*60)))//.replace("T"," "))
        trace.y.push(s.Value)
    })
    console.log(`[tranformSamplesTrace][result] trace sample length:${trace.x.length}; trace name:${trace.name}`);
    console.log(trace);
    return trace
}

console.log("testting transformObjectsTraces__________________")
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

console.log("END_______________________________________________")
function transformObjectsTraces(objects) {
    var traces = []
    trendSelectContent = []
    for (const [key, value] of Object.entries(objects).sort((a, b) => a[1].Id > b[1].Id)) {
        value.Samples.push(value.LastSample)
        traces.push(tranformSamplesTrace(value.Samples,value.Id))

        trendSelectContent.push(value.Id)//test
      }
      console.log(`[transformObjectsTraces][result] traces.length:${traces.length};`);
    return traces
}
console.log("testing transformIds1_____________________________")
console.log(transformIds1("test/t/test/23"))
console.log("END_______________________________________________")
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


