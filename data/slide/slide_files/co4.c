// +build OMIT

#define CR_BEGIN static int state=0; switch(state) { case 0:
#define CR_RETURN(i,x) do { state=__LINE__; return x; case __LINE__:; } while (0)
#define CR_FINISH }

int function(void) {
	static int i;
	CR_BEGIN;
	for (i = 0; i < 10; i++)
		CR_RETURN(1, i);
	CR_FINISH;
}

// STOP OMIT
