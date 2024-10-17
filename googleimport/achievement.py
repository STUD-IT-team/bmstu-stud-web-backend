

class Achievement:
    def __init__(self, count : str, description : str):
        self.count = count
        self.description = description
    
    def ToDict(self):
        return {
            'count': self.count,
            'description': self.description
        }
        
        