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