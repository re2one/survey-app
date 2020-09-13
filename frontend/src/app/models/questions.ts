import {Survey} from './survey';

export class Question{
  public ID: number;
  public CreatedAt: string;
  public UpdatedAt: string;
  public DeletedAt: string;
  public Survey: Survey;
  public surveyid: number;
  public title: string;
  public text: string;
  public type: string;
}

export class Questions {
  public questions: Array<Question>;
}

export class QuestionsResponse {
  public question: Question;
}
