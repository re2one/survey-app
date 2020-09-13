import {Question} from './questions';

export class Mucho{
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Question: Question;
  public questionId: number;
  public text: string;
}

export class Questions {
  public choices: Array<Mucho>;
}

export class SurveyResponse {
  public question: Mucho;
}
