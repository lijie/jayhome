MDFILES := $(wildcard *.md)
TXTFILES := $(patsubst %.md,%.txt,$(MDFILES))

all: $(TXTFILES)

$(TXTFILES):%.txt:%.md
	markdown $< > $@