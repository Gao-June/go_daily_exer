/*
    我用 cpp 也试了下
*/

#include <iostream>
#include <thread>

int a = 0, b = 0;
bool Flag_global = true;

void Add_nums() {
	for(int i = 1; i < 1000000; i++) {
		a++;
		b++;
	}
    Flag_global = false;
}

void Is_equal() {
	/*
		a_value = a
		b_value = b
	*/
	while( Flag_global ){
		if(a < b) {
			std::cout << "a = " << a <<  "\tb = " <<  b << std::endl;
		}
	}

}

int main( ){
    std::thread Thread_1(Add_nums);
    std::thread Thread_2(Is_equal);

    Thread_1.join();
    Thread_2.join();

    return 0;
}

