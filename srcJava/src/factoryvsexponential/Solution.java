package factoryvsexponential;

import java.util.*;

// https://www.codingame.com/training/hard/factorial-vs-exponential

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
class Solution {

    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int K = in.nextInt();
        double[] values = new double[K];
        double valueTmp;
        float factorielTmp;
        float puissanceTmp;
        StringBuilder result = new StringBuilder();

        for (int i = 0; i < K; i++) {
            float A = in.nextFloat();
            values[i] = (double) A;
        }

        for (int i = 0; i < values.length; i++) {
            valueTmp = values[i];

            System.err.println("\nTraitement : " + valueTmp);

            float cpt = 0;
            while (cpt != -1) {

                factorielTmp = calculFactoriel(cpt);
                puissanceTmp = calculPuissance(valueTmp, (double) cpt);

                System.err.println("\t" + puissanceTmp + " < " + factorielTmp);
                if (puissanceTmp < factorielTmp) {
                    System.err.println("find : " + (int) cpt);
                    result.append((int) cpt);
                    break;
                }
                cpt ++;
            }
            result.append(" ");
        }

        System.out.println(result.toString().trim());

    }

    public static int calculPuissance (double value, double puissance) {
        int sum = 0;

        sum = (int) Math.pow(value, puissance);

        return sum;

    }

    public static float calculFactoriel(float value) {
        if (value == 0) {
            return(1);
        } else {
            return(value * calculFactoriel(value-1));
        }
    }
}