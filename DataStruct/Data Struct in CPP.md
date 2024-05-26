[TOC]

# 数据结构算法母版

## 一.常见常量定义

主要包括 ```Status```的值 和 数据结构内**数据类型**的定义。 

```cpp
#define TRUE 1
#define FALSE 0
#define OK 1
#define ERROR 0 	//数值不合法返回ERROR
#define OVERFLOW -2  //数值溢出时返回OVERFLOW
const int MAXSIZE=0x3f3f3f;
typedef int Status;	//用于特指函数返回值
typedef int ElemType;	//线性表内数据的数据类型
```

## 二.线性表和链表

### 1.线性表

定义线性表：

```c++
typedef struct{
    ElemType *Elem;
    int length;
}SqList;
```

初始化：

```cpp
Status InitList_Sq(SqList &L){
    L.Elem=new ElemType[MAXSIZE];
    if(!L.Elem)exit(OVERFLOW);  //若内存分配失败，则返回OVERFLOW
    L.length=0;  //初始化长度为0
    return OK;
}
```

销毁线性表：

```cpp
void DestroyList(SqList &L)
{
    if(L.Elem)delete L.Elem;
}
```

清空线性表：

```cpp
void ClearList(SqList &L){
    L.length = 0;
}
```

按序查找：

```cpp
Status GetElem(SqList &L,int i,ElemType &e)
{
    if(i<1&&i>L.length)return ERROR;//判断合法性
    e=L.elem[i-1];
    return OK;
}
```

按值查找：

```cpp
int LocateElem(SqList &L,ElemType e){
    int i=0;
    while(i<L.length&&L.Elem[i]!=e)i++;
    if(i<L.length)return i+1;
    return 0;
}
```

插入元素：

在第i个位置插入元素

```cpp
Status ListInsert(SqList &L,int i.ElemType e){
    if(i>L.length+1||i<1)return ERROR;  //判断数值合法性
    if(length==MAXSIZE)return ERROR;  //判断是否可插入数值，即顺序表是否已满
    for(int j=L.length-1;j>=i-1;j--)
        L.Elem[j+1]=L.Elem[j];  //i及后边所以的元素后移
    L.Elem[i-1]=e;
    L.length++;
    return OK;
}
```

删除元素：

在第i个位置删除元素

```cpp
Status ListDelete(SqList &L,int i)
{
    if(i<1||i>L.length)return ERROR;//判断合法性
    for(int j=i;j<=L.length-1;j++)
        L.Elem[j-1]=L.Elem[j];//i后的所有元素前移
    L.length--;
    return OK;
}
```

### 2.链表

#### a.单链表

定义链表：

```cpp
typedef struct LNode{
    ElemType data;
    struct LNode *next;
}LNode,*LinkList;
```

初始化链表：

```cpp
Status InitList(LinkList &L){
    L=new LNode;
    L->next=NULL;
    return OK;
}
```

判断空链表：

```cpp
bool IsEmpty(LinkList L){
	if(L->next)return false;
    else return true;
}
```

链表的销毁：

```cpp
Status DestroyList(LinkList &L){
    LNode *p;//或LinkList p;
    while(L){
        p=L;	//p是中间指针变量，用于删除L的原地址
        L=L->next;
        delete p;
    }
    return OK;
}
```

清空链表：

```cpp
Status ClearList(LinkList &L){
    LNode *p,*q;//要保留L指针，所以需要另置一个中间变量
    p=L->next;//p指向首元结点
    while(p){
        q=p->next;//q指向p的下一位
        delete p;//删除p所指向的地址
        p=q;//p指向下一位
    }
    L->next=NULL;
    return OK;
} 
```

求链表表长：

```cpp
int ListLength(LinkList &L){
    LNode *p;
    int i=0;
    while(p){
        i++;
        p=p->next;
    }
    return i;
}
```

按序查找：

求第i个节点的值

```cpp
int SearchList(LinkList &L,int i，ElemType &e){
    LNode *p;
    p=L->next;
    int j=1;
    //此处定义一个新计数器的目的：验证i的合法性，防止i不合法导致程序错误
    while(p&&j<i){
       	p=p->next;
        ++j;	//和j++一个意思
    }
    if(!p||j>i)return ERROR; //第i个元素不存在；
    e=p->data; //将第i个元素的值赋给e
    return OK;
}
```

按值查找：

求与e相等的节点的序号

```cpp
int LocateList(LinkList &L,ElemType e){
    LNode *p;
    int j=1;
    p=L->next;
    while(p&&p->data!=e){
        p=p->next;
        ++j;
    }
    if(!p)return NULL; //若无该元素，则返回空(NULL)
    return j;
}
```

插入元素：

在第i个节点后插入一个新的节点

```cpp
Status InsertList(LinkList &L,int i,ElemType e){
    LNode *p,*s;
    int j=0;
    p=L;
    while(p&&i-1>j){
        p=p->next;
        ++j;
    }		//此时若查找到，则p是指向第i-1个节点
    if(!p||j>i-1)return ERROR; //验证合法性
    s=new LNode;
    s->data=e;
    
    s->next=p->next;
    p->next=s;	//指针的转接，令p指向s，s指向原本p的指向
    
    return OK;
}
```

删除元素：

删除第i个节点

```cpp
Status RemoveList(LinkList &L,int i,ElemType &e){
    LNode *p,*q;
    int j=1;
    p=L->next;
    while(p->next&&j<i-1){
        p=p->next;
        ++j;
    }
    if(!(p->next)||j>i-1)return ERROR;  //判断合法性
    
    q=p->next; //令q指向第i个节点(即要删除的节点)
    p->next=q->next; //第i-1个节点指向第i+1个节点
    e=q->data; //取出第i个节点的数值
    delete q; //释放q的空间
    
    return OK;
}
```

头插法建立链表：

```cpp
void HeadCreate(LinkList &L,int n)//n是链表的长度
{
    LNode *p;
    L=new LNode;
    L->next=NULL;
    while(n--){
        p=new LNode;
        cin >> p->data;
        p->next=L->next;
        L->next=p;
    }
}
```

尾插法建立链表：

```cpp
void TailCreate(LinkList &L,int n){
    LNode *p,*r;
    L=new LNode;
    L->next=NULL;
    r=L;
    while(n--){
        p=new LNode;
        cin >> p->data;
        p->next=NULL;
       	r->next=p;
        r=r->next;//即r=p;
    }
}
```

#### b.双链表

定义双链表：

```cpp
typedef struct DBLNode{
    ElemType data;
    struct DBLNode *prior,*next;
}DBLNode,*DBLinkList;
```

插入元素：

在第i个节点前插入新节点

```cpp
Status InsertDBList(DBLinkList &L,ElemType e,int i){
    LNode *p,*s;
    int j=1;
    p=L->next;
    while(p&&i>j){
        p=p->next;
        ++j;
    }
    if(!p||j>i)return ERROR;
    
    s=new DBLNode;
    s->data=e;
    s->prior=p->prior;
    p->prior->next=s;
    s->next=p;
    p->prior=s;
    
    return OK;
}
```

删除元素：

删除第i个节点

```cpp
Status DeleteDBList(DBLinkList &L,int i,ElemType &e){
    LNode *p;
    p=L->next;
    int j=1;
    while(p&&i>j)j++;
    if(!p||i<j)return ERROR;
    
    e=p->data;
    p->prior->next=p->next;
    p->next->prior=p->prior;
    delete p;
    return OK;
}
```

## 三.栈与队列

### 1.栈

#### a.顺序栈

定义顺序栈：

```cpp
typedef int SElemType;
typedef struct{
    SElemType *base;
    SElemType *top;
    int StackSize;
}SqStack;
```

初始化：

```cpp
Status InitStack(SqStack &S){
    S.base=new SElemType[MAXSIZE];
    if(!S.base) exit(OVERFLOW);
    S.top=S.base;
    S.StackSize=MAXSIZE;
    return OK;
}
```

判断空栈：

```cpp
Status IsEmpty(SqSatack &S){
    if(S.top==S.base)return TRUE;
    else return FALSE;
}
```

清空顺序栈：

```cpp
Status ClearStack(SqStack &S){
    if(S.base)S.top=S.base;
    return OK;
}
```

销毁顺序栈：

```cpp
Status DestroyStack(SqStack &S){
    if(S.base){
        delete S.base; 	//内存被清除了，但指针还存在
        S.StackSize=0;
        S.base=S.top=NULL; 	//使指针指向NULL，防止出现野指针
    }
    return OK;
}
```

入栈：

```cpp
Status Push(SqStack &S,SElemType e){
    if(S.top-S.base==S.StackSize)return ERROR;
    *S.top++=e;
    return OK;
}
```

出栈：

```cpp
Status Pop(SqStack &S,SElemType e){
    if(S.top==S.base)return ERROR;
    e=*--S.top;
    return OK;
}
```

#### b.链栈

定义链栈：

```cpp
typedef struct StackNode{
    SElemType data;
    struct StackNode *next;
}StackNode,*LinkStack;
```

初始化：

```cpp
Status StackEmpty(LinkStack &S){
    S=NULL;
    return OK;
}
```

判断空栈：

```cpp
Status IsEmpty(StackLink &S){
    if(S==NULL)return TRUE;
    else return FALSE;
}
```

入栈：

```cpp
Status Push(LinkStatus &S,SElemType e){
    p=new StackNode;
    p->data=e;
    p->next=S;
    S=p;
    return OK;
}//S为栈顶的指针
```

出栈：

```cpp
Status Pop(LinkStack &S,SElemType &e){
    if(!S)return ERROR; //判断是否为空栈
    e=S->data; 	//取出栈顶的值
    p=S;	
    S=S->next; //S下移
    delete p; 	//删除p指针，由于p为函数内变量，无需改为NULL
    return OK;
}
```

取栈顶元素：

```cpp
SElemType Gettop(LinkStack &S){
    if(S)return S->data;
}
```

### 2.队列

#### a.顺序循环队列

定义顺序循环队列：

```cpp
#define MAXSIZE 100
typedef struct{
    QElemType *base;
    int front; 	//头指针
    int rear;	//尾指针
}SqQueue;
```

初始化：

```cpp
Status InitQueue(SqQueue &Q){
    Q.base=new QElemType[MAXSIZE];
    //Q.base=(QElemType*)malloc(MAXSIZE*sizeof(QElemType));
    if(!Q.base)exit(OVERFLOW);
    Q.front=Q.rear=0;
    return OK;
}
```

求队列长度：

```cpp
int QueueLength(SqQueue Q){
    return ((Q.rear-Q.front+MAXSIZE)%MAXSIZE);
}
```

入队：

```cpp
Status EnQueue(SqQueue &Q,QElemType e){
    if((Q.rear+1)%MAXSIZE==Q.front)return ERROR;
    //判断队列是否为队满状态
    Q.base[Q.rear]=e;
    Q.rear=(Q.rear+1)%MAXSIZE;
    return OK;
}
```

出队：

```cpp
Status DeQueue(SqQueue &Q,QElemType &e){
    if(Q.rear==Q.front)return ERROR;//判断队空
    e=Q.base[Q.front];
    Q.front=(Q.front+1)%MAXSIZE;
    return OK;
}
```

#### b.链式队列

定义链式队列：

```cpp
#define MAXSIZE 100
typedef struct QNode{
    QElemType data;
    struct QNode *next;
}QNode,*QueuePtr;

typedef struct{
    QueuePtr front;
    QueuePtr rear;
}LinkQueue;
```

初始化：

```cpp
Status InitQueue(LinkQueue &Q){
    Q.front=Q.rear=new QueuePtr;
    if(!Q.front)exit(OVERFLOW);
    Q.front->next=NULL;
    return OK;
}
```

销毁链式队列：

```cpp
Status DestroyQueue(LinkQueue &Q){
    QueuePtr p;
    while(Q.front){
        p=Q.front->next;
        delete Q.front;
        Q.front=p;
    }//存在循环，若Q.front不为空，则删除Q.front指向的QNode
    return OK;
}
```

入队：

```c++
Status EnQueue(QNode &Q,QElemType e){
    QueuePtr p;
    p=new QNode;
    if(!p)exit(OVERFLOW);//判断是否成功分配
   	p->data=e;
    p->next=NULL;//队尾插入，则next指向NULL
    Q.rear->next=p;
    //Q.rear此时指向p的前一个节点，即让前一个节点的next指向p
    Q.rear=p;//将Q.rear指向p
    return OK;
}
```

出队：

```c++
Status DeQueue(QNode &Q,QElemType &e){
    QueuePtr p;
    if(Q.front==Q.rear)return ERROR;
    p=Q.front->next;//首元节点是Q.front->next
    e=p->data;
    Q.front->next=p->next;
    if(Q.rear==p)Q.rear=Q.front; 
    delete p;
    return OK;
}
```

## 三.串、数组和广义表

### 1.串

#### a.顺序串

定义顺序串：

```cpp
typedef struct{
    char ch[MAXLEN+1];
    int length;
}SqString;
```

#### b.链串

定义链串：

```cpp
#define CHUNKSIZE 80
typedef struct Chunk{
    char ch[CHUNKSIZE];
    struct Chunk *next;
}Chunk;

typedef struct{
    Chunk *head,*tail; //串的头指针和尾指针
    int curlen; //串的当前长度
}LString; //块链结构
```

#### c.模式匹配算法

##### I.BF算法

```cpp
int BFIndex(SqString S,SqString T){
    int j=1,i=1;
    while(i<=S.length&&j<=T.length){
        if(S.ch[i]==T.ch[i]){
            ++i;
            ++j;
        }
        else{
            i=i-j+2;//回溯
            j=1;
        }
    }
    if(j>=T.length)return i-T.length;//返回第一个匹配的字符的下标
    else return 0;
}
```

##### II.KMP算法

**next[j]**的取值：

1.当j=1时，next[j]=0

2.当子串的前k-1个元素与P[j]前的k-1个元素完全相等时，next[j]取k的最大值

3.若不满足上方的情况时，next[j]=1

```c++
void getNext(SqString T,int &next[]){
    int i=1,next[1]=0;j=0;
    while(i<T.length){
        if(j==0||T.ch[i]==T.ch[j]){
            ++i;
            ++j;
            next[i]=j;
        }
        else
            j=next[j];
    }
}
```

有了next数组后的KMP算法：

```c++
int KMPIndex(SqString S,SqString T){
    int j=1,i=1;
    while(i<=S.length&&j<=T.length){
        if(j==0||S.ch[i]==T.ch[i]){
            i++;
            j++;
        }
        else{
            j=next[j];//i不回溯，j去next[j]的值
        }
        if(j>T.length)return i-T.length;
        else return 0;
    }
}
```

改进next数组得到nextval[]，关键在于判断next[j]是否与j指向的字符一致，若一致则再去next[next[j]]，直至不同或回到串首。

```c++
void getNextval(SqString T,int &nextval[]){
    int i=1,nextval[1]=0,j=0;
    while(i<T.length){
        if(j==0||T.ch[i]==T.ch[j]){
            i++;
            j++;
            if(T.ch[i]!=T.ch[j])nextval[i]=j;
            else nextval[i]=nextval[j];
        }
        else j=nextval[j];
    }
}
```

### 2.数组

稀疏矩阵的压缩方法：

#### a.三元组法

M由三元组(i,j,a~ij~)唯一确定矩阵的非零单元，三元组和矩阵维数可唯一确定一个稀疏矩阵。

```c++
typedef struct trimat{
    ElemType data;
    int i;
    int j;
    struct trimat *next;
}trimat;
```

#### b.十字链表法

结构体中除了安排三元组的三个元素外，还有right和down

其中**right**用于链接同一行中下一个非零元素

其中**down**用于链接同一列中下一个非零元素

```c++
typedef struct trimat{
    ElemType data;
    int i;
    int j;
    struct trimat *right;
    struct trimat *down;
}trimat;
```
