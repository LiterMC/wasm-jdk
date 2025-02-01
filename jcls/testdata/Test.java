
public class Test {
	private int privateInt0 = 0;
	int int1 = 1;
	public long publicLong16 = 16;
	public final String publicFinalString = "publicFinalString";

	private void testPrivateVoidMethod() {
		System.out.println("testPrivateVoidMethod");
	}

	public final int testPublicFinalAddMethod(int a, int b) {
		return a + b;
	}

	public static void main(String[] args) {
		System.out.print("arg0: " + args[0] + "\n");
		System.out.println("Test class");
	}
}
