package main

func doSomething() error {
	return nil
}

func doAnotherThing() error {
	return nil
}

// Func1  bad example
func Func1() error {
	err := doSomething()
	if err == nil {
		err := doAnotherThing()
		if err == nil {
			return nil
		}
		return err
	}
	return err
}

// Func2 good example
func Func2() error {
	if err := doSomething(); err != nil {
		return err
	}
	if err := doAnotherThing(); err != nil {
		return err
	}
	return nil
}

func main() {

}
