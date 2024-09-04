#include <iostream>
using namespace std;

int main()
{
    int n, f0 = 0, f1 = 1, berikutnya = 0;

    cout << "Masukan Batas Deret Bilangan Fibonacci :  ";
    cin >> n;
	cout<<endl;
    cout << "Deret Fibonacci: ";

    for (int i = 0; i <= n; ++i)
    {
        // Mencetak dua deret bilangan fibonacci pertama.
        if(i == 0)
        {
            cout << " " << f0<<" ";
            continue;
        }
        if(i == 1)
        {
            cout << f1 << " ";
            continue;
        }
        berikutnya = f0 + f1;
        f0 = f1;
        f1 = berikutnya;
         // Mencetak deret bilangan fibonacci berikutnya.
        cout << berikutnya << " ";
    }
    return 0;
}