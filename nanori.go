package main

type HTMLObject struct {
	children []*HTMLObject
	tagName  string
}

func getElementsByTagName(element *HTMLObject) {
	var elements []*HTMLObject

	var children []*HTMLObject
	children = element.
	for index := 0; index < element
}

func main() {
	var htmlObject HTMLObject
	htmlObject.tagName = "TBODY"
	getElementsByTagName(&htmlObject)
}

/*
function getElementsByTagName(element, tagName) {
	// Empty array so we can fill it
	var elements = []

	element.children = [
		HTMLOBJECT(TR)
		HTMLOBJECT(TR)
		HTMLOBJECT(TR)
		HTMLOBJECT(TR)
		HTMLOBJECT(TR)
	]

	var children = element.children
	for (var index = 0; index < children.length; index ++) {
		var child = children[index]
		if (child.tagName == tagName) {
			elements.push(child)
		}
	}

	return elements
}
*/
