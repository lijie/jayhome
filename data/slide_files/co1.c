// +build OMIT

void function1() {
	static int step = 0;
	switch (step) {
	case 1:
		dostep1();
		step++;
		break;

	case 2:
		dostep2();
		step++;
		break;	
	}
}

void main() {
	while (1) {
		function1();
		function2();
	}
}

// STOP OMIT
