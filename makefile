.PHONY: update run


update:
	hz update -idl ./idl/picture.thrift

run:
	sh build.sh && sh output/bootstrap.sh