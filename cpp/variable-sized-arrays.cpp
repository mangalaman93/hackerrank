#include <cstdio>
using namespace std;

int main() {
    int N, Q, k;
    int **data;
    scanf("%d %d", &N, &Q);
    data = new int*[N];
    for(int i=0; i<N; i++) {
        scanf("%d", &k);
        data[i] = new int[k];
        for(int j=0; j<k; j++) {
            scanf("%d", &data[i][j]);
        }
    }

    int a, b;
    for(int i=0; i<Q; i++) {
        scanf("%d %d", &a, &b);
        printf("%d\n", data[a][b]);
    }

	return 0;
}
