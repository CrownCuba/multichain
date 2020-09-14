import sys
import os


def specForConfFiles(directoryList: list) -> list:
    confFiles = []
    for direc in directoryList:
        currentDir = pathSandbox + direc +'/'
        confFiles += [currentDir+ x for x in filter(lambda x: os.path.isfile(
            currentDir + x ) and x[-5:] == '.conf', os.listdir( currentDir ))]
    return confFiles

def writeInfo(confFiles: list, infoForAppend: list) -> None:
    for f in confFiles:
        fd = open(f, 'a')
        for line in infoForAppend:
            fd.write(line)
            fd.write('\n')
        fd.write('\n')
        fd.close()

if __name__ == "__main__":
    pathSandbox = sys.argv[1]
    infoForAppend = ['rpcuser=user', 'rpcallowip=0.0.0.0/0']
    directoryList = [x for x in filter(lambda x: os.path.isdir(pathSandbox + x),
    os.listdir(pathSandbox))]
    confFiles = specForConfFiles(directoryList)
    writeInfo(confFiles, infoForAppend)
    
