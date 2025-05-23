#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <set>
#include <algorithm>

using namespace std;

int main() {
    // Чтение чисел из файлов A.txt и B.txt
    ifstream fileA("A.txt");
    ifstream fileB("B.txt");
    
    vector<int> numbersA, numbersB;
    int num;
    
    while (fileA >> num) numbersA.push_back(num);
    while (fileB >> num) numbersB.push_back(num);
    
    fileA.close();
    fileB.close();
  
    return 0;
}
