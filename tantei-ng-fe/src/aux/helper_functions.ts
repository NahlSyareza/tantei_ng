// Min and max number is included as a possibilty
export function getRandomNum(min: number, max: number): number {
	return Math.floor(Math.random() * (max - min + 1)) + min;
}

// Yo credits yo Fisher-Yates man. Cool algorithm
export function fisherYatesShuffler<T>(items: T[]) {
	// let sentinel : T[] = [];

	let lastIndex: number = items.length - 1;

	while (lastIndex > 1) {
		const diceNum: number = getRandomNum(0, lastIndex);
		if (diceNum != lastIndex) {
			const temp: T = items[diceNum];
			items[diceNum] = items[lastIndex];
			items[lastIndex] = temp;
		}
		lastIndex--;
	}
}
