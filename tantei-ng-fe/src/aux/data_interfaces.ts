export interface Studyword {
	kanji: string;
	furigana: string;
	latin: string;
	english: string;
	radical: string;
}

export interface Studyset {
	_id: string;
	name: string;
	items: Studyword[];
	// owner: string;
	indexed_radicals: string[];
}

export interface RadicalList {
	_id: string;
	base_radical: string;
	radical_variants: string[];
}
