
def ParseSharedFolderID(url: str) -> str:
    """Parses the Shared Folder ID from the Google Drive URL."""
    folderIndex = 0
    splitURL = url.split('/')
    while folderIndex < len(splitURL) - 1:
        if splitURL[folderIndex] == 'folders':
            return splitURL[folderIndex + 1]
        folderIndex += 1
    raise ValueError('Could not parse Shared Folder ID: no folders found')
    
def ParseSharedFileID(url: str) -> str:
    """Parses the Shared File ID from the Google Drive URL."""
    fileIndex = 0
    splitURL = url.split('/')
    while fileIndex < len(splitURL) - 1:
        if splitURL[fileIndex] == 'file':
            return splitURL[fileIndex + 2]
        fileIndex += 1
    raise ValueError('Could not parse Shared File ID: no files found')

