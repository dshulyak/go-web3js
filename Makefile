default:
	npm install
	cp ./node_modules/web3/dist/web3.min.js web3.js
	go run cmd/bundler.go --src web3.js -dst web3js.go
	rm web3.js
