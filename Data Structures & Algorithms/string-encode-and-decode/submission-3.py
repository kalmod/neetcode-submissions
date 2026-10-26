class Solution:

    def encode(self, strs: List[str]) -> str:
        encodedStr = ""
        for s in strs:
            encodedStr += str(len(s)) + "#" + s
        return encodedStr


    def decode(self, s: str) -> List[str]:
        decodedStrs = []
        i = 0
        while i < len(s):
            # A delimiter is need to signify the end of the integer characters representing the length
            lengthOfStr = ""
            while s[i] != "#":
                lengthOfStr += s[i]
                i += 1
            i += 1  # Skip the '#'
            length = int(lengthOfStr)
            decodedStrs.append(s[i:i+length])
            i += length
        return decodedStrs
