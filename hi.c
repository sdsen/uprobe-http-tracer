#include<stdio.h>
 int check(char* ip, int v){
	 printf("%s\n",ip);
	 return v;
 }
int main(){
	printf("Hello\n");
	check("Hi", 10);
}
