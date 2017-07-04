diagrams:
	go test -tags diagrams -run TestDiagrams

clean:
	$(RM) *.png
