func Test#[[$NAME$]]#(${PARAM_NAME} ${PARAM_TYPE}) {
	tl := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, tc := range tl {
		${PARAM_NAME}.Run(tc.name, func(${PARAM_NAME} ${PARAM_TYPE}) {
			#[[$END$]]#
		})
	}
}
