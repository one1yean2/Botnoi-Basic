import pandas as pd

FILENAME = 'grades_with_grades.xlsx'

def excelGradeCalculator(filename):
    
    #read from excel
    df = pd.read_excel(filename)
    #loop through every instances
    for inx,row in df.iterrows():
        #update grade 
        df['เกรด'][inx] = gradeCalculator(row['คะแนน'])
    #sort by grade
    df.sort_values(by=['เกรด'],inplace=True)
    #writing to excel
    df.to_excel(filename,index=False)
    # df.to_excel("result.xlsx",index=False)
    
def gradeCalculator(score):
    
    #check type of score
    if(type(score) is int):
        if(score >= 80):
            return "A"
        elif(score >= 70):
            return "B"
        elif(score >= 60):
            return "C"
        elif(score >= 50):
            return "D"
        else:
            return "F"
    else:
            return "F"
        
excelGradeCalculator(FILENAME)
