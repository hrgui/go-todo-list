export interface Todo {
  id: number;
  title: string;
  completed: boolean;
  created_date?: Date;
  updated_date?: Date;
}

export type TodoComponentProps = Partial<Todo>;
