// Code generated by smithy-go-codegen DO NOT EDIT.
package types

type ComplexNestedErrorData struct {
	Foo *string
}

type NestedPayload struct {
	Greeting *string
	Name     *string
}

type RecursiveShapesInputOutputNested1 struct {
	Foo    *string
	Nested *RecursiveShapesInputOutputNested2
}

type RecursiveShapesInputOutputNested2 struct {
	Bar             *string
	RecursiveMember *RecursiveShapesInputOutputNested1
}

type StructureListMember struct {
	A *string
	B *string
}

type GreetingStruct struct {
	Hi *string
}
