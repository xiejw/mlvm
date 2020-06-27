TEX=docker run --rm -v `pwd`:/workdir xiejw/tex pdftex

default: 01-tensor.pdf 02-conv2d.pdf

01-tensor.pdf: 01-tensor.tex
	${TEX} 01-tensor.tex

02-conv2d.pdf: 02-conv2d.tex
	${TEX} 02-conv2d.tex

clean:
	rm -rf *.pdf *.log
