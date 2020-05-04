TEX=docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

01-tensor.pdf: 01-tensor.tex
	${TEX} 01-tensor.tex

clean:
	rm -rf *.pdf *.log
