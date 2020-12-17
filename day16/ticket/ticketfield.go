package ticket

type Field struct {
	RangeLowerFirst, RangeUpperFirst   int
	RangeLowerSecond, RangeUpperSecond int
}

func NewTicketField(rlf, ruf, rls, rus int) *Field {
	return &Field{
		RangeLowerFirst:  rlf,
		RangeUpperFirst:  ruf,
		RangeLowerSecond: rls,
		RangeUpperSecond: rus,
	}
}

func (tf *Field) IsValidField(Field int) bool {
	return (Field >= tf.RangeLowerFirst && Field <= tf.RangeUpperFirst) ||
		(Field >= tf.RangeLowerSecond && Field <= tf.RangeUpperSecond)
}
