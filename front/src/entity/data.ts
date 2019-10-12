export interface Data {
  createdAt: string;
  amount: number;
  diff: number;
  totalAmount: number;
}

export interface DataList {
  data: Data[];
  isLoading: boolean;
}
