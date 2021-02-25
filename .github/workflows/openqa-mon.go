name: openqa-mon
on: push

jobs:
  openqa-mon:
    name: openqa-mon
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
      - name: Setup go
        uses: actions/setup-go@v2
        with:
          go-version: '1.14'
      - name: Install requirements
        run: echo "No requirements"
      - name: Build openqa-mon
        run: make openqa-mon
