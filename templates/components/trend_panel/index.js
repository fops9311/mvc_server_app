
var trendtrace1 = {

    type: "scatter",
  
    mode: "lines",
  
    name: 'fops9311@yandex.ru/v1',
  
    
    x: ['2013-10-04 22:23:00', '2013-11-04 22:23:00', '2013-12-04 22:23:00'],

    y: [1, 3, 6],
  
    line: {color: '#17BECF'}
  
  };
  
  var trendtrace2 = {
  
    type: "scatter",
  
    mode: "lines",
  
    name: 'fops9311@yandex.ru/v2',
  
    
    x: ['2013-10-04 22:23:00', '2013-11-04 22:23:00', '2013-12-04 22:23:00'],

    y: [10, 30, 16],
  
    line: {color: '#7F7F7F'}
  
  };
  
  var trendtrace3 = {
  
    type: "scatter",
  
    mode: "lines",
  
    name: 'fops9311@yandex.ru/v3',
  
    
    x: ['2013-10-04 22:23:00', '2013-11-04 22:23:00', '2013-12-04 22:23:00'],

    y: [101, 20, 56],
  
    line: {color: '#7F007F'}
  
  };
  
  var trendData = [trendtrace1,trendtrace2,trendtrace3];
  
  
  var trendlayout = {
  
    title: 'Time Series with Rangeslider',
  
    xaxis: {
  
      autorange: true,
  
      range: ['2015-02-17', '2017-02-16'],
  
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
  
      rangeslider: {range: ['2015-02-17', '2017-02-16']},
  
      type: 'date'
  
    },
  
    yaxis: {
  
      autorange: true,
  
      range: [86.8700008333, 138.870004167],
  
      type: 'linear'
  
    }
  
  };
  Plotly.newPlot('myTrend', trendData, trendlayout,{responsive: true});
  
  
