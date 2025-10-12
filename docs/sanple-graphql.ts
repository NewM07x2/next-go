export type Maybe<T> = T | null;
export type InputMaybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };

/** GraphQLスキーマに基づく型定義 */
export type User = {
  __typename?: 'User';
  id: string;
  name: string;
  email: string;
};

export type Query = {
  __typename?: 'Query';
  user?: Maybe<User>;
};

export type QueryUserArgs = {
  id: string;
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser?: Maybe<User>;
};

export type MutationCreateUserArgs = {
  name: string;
  email: string;
};

/** クエリに基づく型定義 */
export type GetUserQueryVariables = Exact<{
  id: string;
}>;

export type GetUserQuery = {
  __typename?: 'Query';
  user?: {
    __typename?: 'User';
    id: string;
    name: string;
    email: string;
  } | null;
};

export type CreateUserMutationVariables = Exact<{
  name: string;
  email: string;
}>;

export type CreateUserMutation = {
  __typename?: 'Mutation';
  createUser?: {
    __typename?: 'User';
    id: string;
    name: string;
  } | null;
};