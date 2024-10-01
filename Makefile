run:
	dlv debug  --headless --api-version=2 --listen=127.0.0.1:43000 ./pkg

debug:
	dlv connect 127.0.0.1:43000
