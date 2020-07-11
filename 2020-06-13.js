function testicles(weight) {
    weight = 3
}

var myWeight = 700
testicles(myWeight)
console.log(myWeight)



function changeInnerHTML(element, text) {
    element.innerHTML = text
}

var spanElement = document.getElementById("mySpan")
changeInnerHTML(spanElement, "Hi there")
console.log(spanElement.innerHTML)

var reference1 = spanElement
var reference2 = spanElement
var reference3 = reference2
var reference4 = reference3
var reference85 = reference1
var reference23894791823 = reference3
reference4 = reference85

reference23894791823.innerHTML = "Hey Nanori I'm a reference"

var reference1 = "apple"
var reference2 = "apple"
var reference3 = reference2
var reference4 = reference3
var reference85 = reference1
var reference23894791823 = reference3
reference4 = reference85

reference23894791823 = "not apple"

<span id="mySpan">Hi there</span>



// 700