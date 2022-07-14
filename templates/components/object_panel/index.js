var config = {responsive: true}

//Plotly.newPlot('myDiv', data, layout, config)

var myPlot = document.getElementById('myDiv')

myPlot.on('plotly_click', function(data){
  localStorage.setItem('current_obj', data.points[0].label);
  console.log(localStorage.getItem('current_obj') );
});