function changeHelmetColor(){
  var color = document.getElementById('color').value;
  var helmetElements = document.getElementsByClassName("label-helmet")
  for (let element of helmetElements) {
    element.style.backgroundColor = color; 
  }
}
var alertList = document.querySelectorAll('.alert')
var alerts =  [].slice.call(alertList).map(function (element) {
  return new bootstrap.Alert(element)
})

