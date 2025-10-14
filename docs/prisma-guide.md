# Prisma ガイド

## 概要

Prisma は、Node.js 環境で動作する ORM（Object-Relational Mapping）ツールです。
他の言語（例: Python や Java）では直接使用できませんが、API を介して連携することが可能です。データベース操作を簡素化し、TypeScript と統合することで型安全性を提供します。

### 主な特徴

- **型安全性**: データベーススキーマに基づいた型安全なクエリ。
- **簡単なデータベース操作**: Prisma Client を使用して CRUD 操作を簡単に実行可能。
- **マイグレーション管理**: Prisma Migrate を使用してスキーマのバージョン管理が可能。
- **GraphQL との統合**: GraphQL リゾルバの実装を簡素化。

---

## 導入方法

### 1. Prisma のインストール

以下のコマンドを実行して Prisma をインストールします。

```bash
npm install prisma --save-dev
npm install @prisma/client
```

### 2. Prisma の初期化

以下のコマンドを実行して Prisma を初期化します。

```bash
npx prisma init
```

これにより、以下のファイルが生成されます：

- `prisma/schema.prisma`: データベーススキーマを定義するファイル。
- `.env`: データベース接続情報を設定する環境変数ファイル。

### 3. データベーススキーマの定義

`prisma/schema.prisma` にデータベーススキーマを定義します。以下は例です：

```prisma
generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

model User {
  id    Int     @id @default(autoincrement())
  name  String
  email String  @unique
  posts Post[]
}

model Post {
  id        Int      @id @default(autoincrement())
  title     String
  content   String
  published Boolean  @default(false)
  author    User     @relation(fields: [authorId], references: [id])
  authorId  Int
}
```

### 4. マイグレーションの実行

以下のコマンドを実行してデータベースにスキーマを適用します：

```bash
npx prisma migrate dev --name init
```

### 5. Prisma Client の生成

以下のコマンドを実行して Prisma Client を生成します：

```bash
npx prisma generate
```

---

## 使用方法

### Prisma Client の初期化

`lib/prisma.ts` を作成し、Prisma Client を初期化します：

```typescript
import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

export default prisma;
```

### GraphQL リゾルバでの使用

Prisma Client を使用して GraphQL リゾルバを実装します：

```typescript
import prisma from '../lib/prisma';

export const resolvers = {
  Query: {
    users: async () => {
      return await prisma.user.findMany();
    },
  },
  Mutation: {
    createUser: async (_: any, args: { name: string; email: string }) => {
      return await prisma.user.create({
        data: {
          name: args.name,
          email: args.email,
        },
      });
    },
  },
};
```

---

## サンプルクエリ

### ユーザー一覧の取得

```graphql
query {
  users {
    id
    name
    email
  }
}
```

### ユーザーの作成

```graphql
mutation {
  createUser(name: "John Doe", email: "john.doe@example.com") {
    id
    name
    email
  }
}
```

---

## 参考リンク

- [Prisma 公式ドキュメント](https://www.prisma.io/docs)
- [Prisma GitHub リポジトリ](https://github.com/prisma/prisma)
