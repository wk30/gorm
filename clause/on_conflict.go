package clause

type OnConflict struct {
	Columns   []Column
	Where     Where
	DoNothing bool
	DoUpdates Set
}

func (OnConflict) Name() string {
	return "ON CONFLICT"
}

// Build build onConflict clause
func (onConflict OnConflict) Build(builder Builder) {
	if len(onConflict.Columns) > 0 {
		builder.WriteQuoted(onConflict.Columns) // FIXME columns
		builder.WriteByte(' ')
	}

	if len(onConflict.Where.Exprs) > 0 {
		builder.WriteString("WHERE ")
		onConflict.Where.Build(builder)
		builder.WriteByte(' ')
	}

	if onConflict.DoNothing {
		builder.WriteString("DO NOTHING")
	} else {
		builder.WriteString("DO UPDATE SET ")
		onConflict.DoUpdates.Build(builder)
	}
}

// MergeClause merge onConflict clauses
func (onConflict OnConflict) MergeClause(clause *Clause) {
	clause.Expression = onConflict
}
